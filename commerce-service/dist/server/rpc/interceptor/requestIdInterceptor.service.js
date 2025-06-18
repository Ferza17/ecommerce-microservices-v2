"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.RequestIdInterceptor = void 0;
const common_1 = require("@nestjs/common");
const header_1 = require("../../../enum/header");
const uuid_1 = require("uuid");
let RequestIdInterceptor = class RequestIdInterceptor {
    intercept(context, next) {
        const metadata = context.switchToRpc().getContext();
        let requestId = metadata.get(header_1.Header.X_REQUEST_ID)[0];
        if (!requestId) {
            requestId = (0, uuid_1.v4)();
        }
        metadata.set(header_1.Header.X_REQUEST_ID, requestId);
        return next.handle();
    }
};
exports.RequestIdInterceptor = RequestIdInterceptor;
exports.RequestIdInterceptor = RequestIdInterceptor = __decorate([
    (0, common_1.Injectable)()
], RequestIdInterceptor);
//# sourceMappingURL=requestIdInterceptor.service.js.map