"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var ConsulService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConsulService = void 0;
const common_1 = require("@nestjs/common");
const consul_1 = __importDefault(require("consul"));
const config_1 = require("@nestjs/config");
let ConsulService = ConsulService_1 = class ConsulService {
    configService;
    logger = new common_1.Logger(ConsulService_1.name);
    consul;
    constructor(configService) {
        this.configService = configService;
        this.consul = new consul_1.default({
            host: configService.get('CONSUL_HOST'),
            port: configService.get('CONSUL_PORT'),
        });
    }
    async onModuleInit() {
        await this.consul.agent.service.register({
            name: await this.get('/services/commerce/SERVICE_NAME'),
            port: parseInt(await this.get('/services/commerce/RPC_PORT')),
            address: await this.get('/services/commerce/RPC_HOST'),
            tags: ['v1'],
        });
    }
    async get(key) {
        const k = `${this.configService.get('ENV')}${key}`;
        const pair = await this.consul.kv.get(k);
        const value = pair?.Value;
        if (!value) {
            this.logger.error(`Key ${key} not found`);
            throw new Error(`Key ${key} not found in Consul`);
        }
        return value.toString();
    }
};
exports.ConsulService = ConsulService;
exports.ConsulService = ConsulService = ConsulService_1 = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, common_1.Inject)()),
    __metadata("design:paramtypes", [config_1.ConfigService])
], ConsulService);
//# sourceMappingURL=consul.service.js.map