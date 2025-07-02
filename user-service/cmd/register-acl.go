package cmd

import (
	"context"
	"errors"
	"fmt"
	postgres2 "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"time"
)

var aclCommand = &cobra.Command{
	Use: "acl",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		tx := postgres2.ProvidePostgresInfrastructure().GormDB.Begin()

		// 1. Find IF ROLE EXISTS THEN CONTINUE ELSE CREATE new ROLE
		// 2. SAVE CREATED ID IN MAP[pb.ROLE.string()]ID
		mapRoles := searchAndRegisterRole(ctx, []pb.EnumRole{
			pb.EnumRole_CUSTOMER,
			pb.EnumRole_CUSTOMER_MEMBERSHIP,
			pb.EnumRole_ADMIN,
			pb.EnumRole_SUPER_ADMIN,
		}, tx)

		descriptorFilePath := "descriptor.pb"
		descriptorBytes, err := ioutil.ReadFile(descriptorFilePath)
		if err != nil {
			log.Fatalf("unable to read descriptor %s: %v. make sure file can be access!.", descriptorFilePath, err)
		}

		var fileDescriptorSet descriptorpb.FileDescriptorSet
		if err := proto.Unmarshal(descriptorBytes, &fileDescriptorSet); err != nil {
			log.Fatalf("unable to unmarshal FileDescriptorSet: %v", err)
		}

		files, err := protodesc.NewFiles(&fileDescriptorSet)
		if err != nil {
			log.Fatalf("unable to parse FileDescriptorSet: %v", err)
		}

		allFullMethodNames := make(map[string]bool)
		var privateMethod, publicMethod int

		for _, fileProto := range fileDescriptorSet.GetFile() {
			fd, err := files.FindFileByPath(fileProto.GetName())
			if err != nil {
				log.Printf("Warning: failed to find descriptor file for %s (or file import): %v", fileProto.GetName(), err)
				continue
			}

			services := fd.Services()
			for i := 0; i < services.Len(); i++ {
				service := services.Get(i)

				methods := service.Methods()
				for j := 0; j < methods.Len(); j++ {
					method := methods.Get(j)

					fullMethodName := fmt.Sprintf("/%s/%s", service.FullName(), method.Name())
					if _, exists := allFullMethodNames[fullMethodName]; !exists {
						allFullMethodNames[fullMethodName] = true
						methodOptions := method.Options().(*descriptorpb.MethodOptions)

						// User Options
						// 1. IF IS PUBLIC is FALSE	-> FIND IF ACCESS CONTROL FULL METHOD IS EXISTS THEN CONTINUE ELSE CREATE NEW ACCESS CONTROL FULL METHOD
						// 2. IF IS PUBLIC is TRUE	-> FIND IF EXCLUDE METHOD IS EXIST THEN CONTINUE ELSE CREATE NEW EXCLUDED PATH
						if proto.HasExtension(methodOptions, pb.E_Acl) {
							t := getUserAclOptions(methodOptions)
							if !t.IsPublic {
								for _, role := range t.Roles {
									v, ok := mapRoles[role.String()]
									if ok {
										ac := &orm.AccessControl{
											ID:             fmt.Sprintf("%s:%s", role.String(), uuid.New().String()),
											FullMethodName: fullMethodName,
											HttpUrl:        t.Http.GetUrl(),
											HttpMethod:     t.Http.GetMethod(),
											RoleID:         v,
											CreatedAt:      now,
											UpdatedAt:      now,
										}

										if t.Broker != nil {
											ac.EventType = t.Broker.EventType
										}

										searchAndRegisterAccessControl(ctx, ac, tx)
									}
								}

								privateMethod++
							} else {
								ex := &orm.AccessControlExcluded{
									ID:             fmt.Sprintf("EXCLUDED:%s", uuid.New().String()),
									FullMethodName: fullMethodName,
									HttpUrl:        t.Http.GetUrl(),
									HttpMethod:     t.Http.GetMethod(),
									CreatedAt:      now,
									UpdatedAt:      now,
								}

								if t.Broker != nil {
									ex.EventType = t.Broker.EventType
								}

								searchAndRegisterAccessControlExcluded(ctx, ex, tx)
								publicMethod++
							}
						}

						if len(methodOptions.ProtoReflect().GetUnknown()) > 0 {
							fmt.Println("  Custom/Unknown Options detected")
						}
					}
				}
			}
		}

		// Register Health Check for consul as excluded method to avoid response error by interceptor
		healthFullMethods := map[string]*pb.HTTP{
			grpc_health_v1.Health_Check_FullMethodName: &pb.HTTP{
				Url:    "/v1/user/check",
				Method: "get",
			},
			grpc_health_v1.Health_Watch_FullMethodName: &pb.HTTP{
				Url:    "/watch",
				Method: "get",
			},
			grpc_health_v1.Health_List_FullMethodName: &pb.HTTP{
				Url:    "/list",
				Method: "get",
			},
			"/grpc.reflection.v1alpha.ServerReflection/": {
				Url:    "/grpc.reflection.v1alpha.ServerReflection/",
				Method: "get",
			},
			// SWAGGER API DOCS
			"/docs/v1/user/service.swagger.json": {
				Url:    "/docs/v1/user/service.swagger.json",
				Method: "get",
			},
		}

		for s, http := range healthFullMethods {
			searchAndRegisterAccessControlExcluded(ctx, &orm.AccessControlExcluded{
				ID:             fmt.Sprintf("EXCLUDED:%s", uuid.New().String()),
				FullMethodName: s,
				HttpUrl:        http.GetUrl(),
				HttpMethod:     http.GetMethod(),
				CreatedAt:      now,
				UpdatedAt:      now,
			}, tx)
		}

		fmt.Println("\n-----------------------------------------------------")
		fmt.Printf("Total %d Full Method Names gRPC found.\n", len(allFullMethodNames))
		fmt.Printf("Total %d Private Full Method Names found.\n", privateMethod)
		fmt.Printf("Total %d Public Full Method Names found.\n", publicMethod)
		tx.Commit()
	},
}

func getUserAclOptions(methodOptions *descriptorpb.MethodOptions) *pb.MethodAccessControl {
	if proto.HasExtension(methodOptions, pb.E_Acl) {
		customValue := proto.GetExtension(methodOptions, pb.E_Acl)
		val, err := proto.Marshal(customValue.(proto.Message))
		if err != nil {
			log.Fatalf("failed marshal custom options: %v", err)
		}

		var tt pb.MethodAccessControl
		if err := proto.Unmarshal(val, &tt); err != nil {
			log.Fatalf("failed unmarshal custom options: %v", err)
		}
		return &tt
	}

	return nil
}

// RETURN MAP WITH KEY pb.role.string with value id
func searchAndRegisterRole(ctx context.Context, roles []pb.EnumRole, tx *gorm.DB) map[string]string {
	var (
		r                        = map[string]string{}
		now                      = time.Now()
		rolePostgresqlRepository = rolePostgresqlRepository.ProvideRoleRepository()
	)

	for _, role := range roles {
		existingRole, err := rolePostgresqlRepository.FindRoleByName(ctx, "", role.String(), tx)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			continue
		}

		if existingRole != nil {
			r[role.String()] = existingRole.ID
			continue
		}

		existingRole, err = rolePostgresqlRepository.CreateRole(ctx, "", &orm.Role{
			ID:        fmt.Sprintf("%s:%s", role.String(), uuid.New().String()),
			Role:      role.String(),
			CreatedAt: now,
			UpdatedAt: now,
		}, tx)
		if err != nil {
			continue
		}

		r[role.String()] = existingRole.ID
	}

	return r
}

func searchAndRegisterAccessControl(ctx context.Context, ac *orm.AccessControl, tx *gorm.DB) *orm.AccessControl {
	var (
		accessControlPostgresqlRepository = accessControlPostgresqlRepository.ProvideAccessControlRepository()
	)

	existingAccessControl, err := accessControlPostgresqlRepository.FindAccessControlByRoleIdAndFullMethodName(ctx, "", ac.RoleID, ac.FullMethodName, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if existingAccessControl != nil {
		return existingAccessControl
	}

	existingAccessControl, err = accessControlPostgresqlRepository.CreateAccessControl(ctx, "", ac, tx)
	if err != nil {
		return nil
	}
	return existingAccessControl
}

func searchAndRegisterAccessControlExcluded(ctx context.Context, excluded *orm.AccessControlExcluded, tx *gorm.DB) *orm.AccessControlExcluded {
	var (
		accessControlPostgresqlRepository = accessControlPostgresqlRepository.ProvideAccessControlRepository()
	)

	existingAccessControlExcluded, err := accessControlPostgresqlRepository.FindAccessControlExcludedByFullMethodName(ctx, "", excluded.FullMethodName, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if existingAccessControlExcluded != nil {
		return existingAccessControlExcluded
	}

	existingAccessControlExcluded, err = accessControlPostgresqlRepository.CreateAccessControlExcluded(ctx, "", excluded, tx)
	return existingAccessControlExcluded
}
