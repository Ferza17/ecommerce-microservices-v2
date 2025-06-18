"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RabbitmqConsumer = void 0;
const common_1 = require("@nestjs/common");
const consumer_config_1 = require("./consumer.config");
const core_1 = require("@nestjs/core");
const rabbitmq_module_1 = require("./rabbitmq.module");
const queue_1 = require("../../enum/queue");
class RabbitmqConsumer {
    consulConfig;
    logger = new common_1.Logger(RabbitmqConsumer.name);
    constructor(consulConfig) {
        this.consulConfig = consulConfig;
    }
    async Serve() {
        const options = [
            {
                queue: queue_1.Queue.CART_CREATED,
                option: await new consumer_config_1.RabbitmqOptions(this.consulConfig, queue_1.Queue.CART_CREATED).getRabbitmqOptions(),
            },
            {
                queue: queue_1.Queue.CART_UPDATED,
                option: await new consumer_config_1.RabbitmqOptions(this.consulConfig, queue_1.Queue.CART_UPDATED).getRabbitmqOptions(),
            },
            {
                queue: queue_1.Queue.CART_DELETED,
                option: await new consumer_config_1.RabbitmqOptions(this.consulConfig, queue_1.Queue.CART_DELETED).getRabbitmqOptions(),
            },
        ];
        for (const option of options) {
            this.logger.log(`Rabbitmq Consumer ${option.queue} is running...`);
            const app = await core_1.NestFactory.createMicroservice(rabbitmq_module_1.RabbitmqModule, option.option);
            app.listen();
        }
    }
}
exports.RabbitmqConsumer = RabbitmqConsumer;
//# sourceMappingURL=rabbitmq.js.map