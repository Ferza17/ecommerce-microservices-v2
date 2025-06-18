"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProductServiceCircuitOptions = exports.UserServiceCircuitOptions = void 0;
exports.UserServiceCircuitOptions = {
    timeout: 2000,
    errorThresholdPercentage: 50,
    resetTimeout: 10000,
    name: 'UserRpcService',
    volumeThreshold: 5,
    rollingCountTimeout: 10000,
    rollingCountBuckets: 10,
};
exports.ProductServiceCircuitOptions = {
    timeout: 2000,
    errorThresholdPercentage: 50,
    resetTimeout: 10000,
    name: 'ProductRpcService',
    volumeThreshold: 5,
    rollingCountTimeout: 10000,
    rollingCountBuckets: 10,
};
//# sourceMappingURL=circuitOptions.js.map