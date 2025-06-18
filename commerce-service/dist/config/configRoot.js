"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.RabbitMQRootAsync = exports.MongooseRootAsync = exports.configRoot = void 0;
const joi = __importStar(require("joi"));
const config_1 = require("@nestjs/config");
const environment_1 = __importDefault(require("../enum/environment"));
const mongoose_1 = require("@nestjs/mongoose");
const consul_module_1 = require("./consul.module");
const consul_service_1 = require("./consul.service");
const nestjs_rabbitmq_1 = require("@golevelup/nestjs-rabbitmq");
const exchange_1 = require("../enum/exchange");
exports.configRoot = config_1.ConfigModule.forRoot({
    envFilePath: '.env',
    isGlobal: true,
    validationSchema: joi.object({
        ENV: joi.string()
            .valid(environment_1.default.DEVELOPMENT, environment_1.default.LOCAL, environment_1.default.PRODUCTION)
            .default(environment_1.default.LOCAL),
        CONSUL_HOST: joi.string(),
        CONSUL_PORT: joi.number(),
    }),
});
exports.MongooseRootAsync = mongoose_1.MongooseModule.forRootAsync({
    imports: [consul_module_1.ConsulModule],
    inject: [consul_service_1.ConsulService],
    useFactory: async (configService) => ({
        uri: `mongodb://${await configService.get('/database/mongodb/MONGO_USERNAME')}:${await configService.get('/database/mongodb/MONGO_PASSWORD')}@${await configService.get('/database/mongodb/MONGO_HOST')}:${await configService.get('/database/mongodb/MONGO_PORT')}/${await configService.get('/database/mongodb/MONGO_DATABASE_NAME/COMMERCE')}?authSource=admin`,
    }),
});
exports.RabbitMQRootAsync = nestjs_rabbitmq_1.RabbitMQModule.forRootAsync({
    imports: [consul_module_1.ConsulModule],
    inject: [consul_service_1.ConsulService],
    useFactory: async (configService) => ({
        uri: `amqp://${await configService.get("/broker/rabbitmq/RABBITMQ_USERNAME")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PASSWORD")}@${await configService.get("/broker/rabbitmq/RABBITMQ_HOST")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PORT")}`,
        exchanges: [
            {
                name: exchange_1.Exchange.EventExchange.toString(),
                type: 'direct',
            },
        ],
    }),
});
//# sourceMappingURL=configRoot.js.map