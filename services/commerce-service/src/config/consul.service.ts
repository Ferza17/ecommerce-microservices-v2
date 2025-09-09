import {Inject, Injectable, Logger, OnModuleInit} from "@nestjs/common";
import Consul from "consul";
import {ConfigService} from "@nestjs/config";

@Injectable()
export class ConsulService implements OnModuleInit {
    private readonly logger = new Logger(ConsulService.name);
    private consul: Consul;

    // RABBITMQ CONSUL CONFIG
    RabbitMQUsername: string;
    RabbitMQPassword: string;
    RabbitMQHost: string;
    RabbitMQPort: string;

    // COMMERCE SERVICE
    CommerceServiceServiceName: string;
    CommerceServiceRpcHost: string;
    CommerceServiceRpcPort: string;
    CommerceServiceHttpHost: string;
    CommerceServiceHttpPort: string;

    // USER SERVICE
    UserServiceServiceName: string;
    UserServiceRpcHost: string;
    UserServiceRpcPort: string;
    UserServiceHttpHost: string;
    UserServiceHttpPort: string;

    // PRODUCT SERVICE
    ProductServiceServiceName: string;
    ProductServiceRpcHost: string;
    ProductServiceRpcPort: string;
    ProductServiceHttpHost: string;
    ProductServiceHttpPort: string;

    constructor(@Inject() private readonly configService: ConfigService) {
        this.consul = new Consul({
            host: configService.get<string>("CONSUL_HOST"),
            port: configService.get<number>("CONSUL_PORT"),
        });
    }

    async onModuleInit() {
        // RabbitMQ Config
        this.RabbitMQUsername = await this.get(
            "/broker/rabbitmq/RABBITMQ_USERNAME",
        );
        this.RabbitMQPassword = await this.get(
            "/broker/rabbitmq/RABBITMQ_PASSWORD",
        );
        this.RabbitMQHost = await this.get("/broker/rabbitmq/RABBITMQ_HOST");
        this.RabbitMQPort = await this.get("/broker/rabbitmq/RABBITMQ_PORT");

        // Commerce Service Config
        this.CommerceServiceServiceName = await this.get("/services/commerce/SERVICE_NAME");
        this.CommerceServiceRpcHost = await this.get("/services/commerce/RPC_HOST");
        this.CommerceServiceRpcPort = await this.get("/services/commerce/RPC_PORT");
        this.CommerceServiceHttpHost = await this.get("/services/commerce/HTTP_HOST");
        this.CommerceServiceHttpPort = await this.get("/services/commerce/HTTP_PORT");

        // User Service Config
        this.UserServiceServiceName = await this.get("/services/user/SERVICE_NAME");
        this.UserServiceRpcHost = await this.get("/services/user/RPC_HOST");
        this.UserServiceRpcPort = await this.get("/services/user/RPC_PORT");
        this.UserServiceHttpHost = await this.get("/services/user/HTTP_HOST");
        this.UserServiceHttpPort = await this.get("/services/user/HTTP_PORT");

        // Product Service Config
        this.ProductServiceServiceName = await this.get("/services/product/SERVICE_NAME");
        this.ProductServiceRpcHost = await this.get("/services/product/RPC_HOST");
        this.ProductServiceRpcPort = await this.get("/services/product/RPC_PORT");
        this.ProductServiceHttpHost = await this.get("/services/product/HTTP_HOST");
        this.ProductServiceHttpPort = await this.get("/services/product/HTTP_PORT");

        await this.consul.agent.service.register({
            name: await this.get("/services/commerce/SERVICE_NAME"),
            port: parseInt(await this.get("/services/commerce/RPC_PORT")),
            address: await this.get("/services/commerce/RPC_HOST"),
            tags: ["service", "rabbitmq_client", "rpc", "http"],
        });
    }

    async get(key: string): Promise<string> {
        const k = `${this.configService.get<string>("ENV")}${key}`;
        const pair = await this.consul.kv.get(k);
        const value = pair?.Value;
        if (!value) {
            this.logger.error(`Key ${key} not found`);
            throw new Error(`Key ${key} not found in Consul`);
        }
        return value.toString();
    }
}
