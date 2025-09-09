import { Options as CircuitBreakerOptions } from "opossum";

export const UserServiceCircuitOptions: CircuitBreakerOptions = {
  timeout: 2000,
  errorThresholdPercentage: 50,
  resetTimeout: 10000,
  name: "UserRpcService",
  volumeThreshold: 5,
  rollingCountTimeout: 10000,
  rollingCountBuckets: 10,
};

export const ProductServiceCircuitOptions: CircuitBreakerOptions = {
  timeout: 2000,
  errorThresholdPercentage: 50,
  resetTimeout: 10000,
  name: "ProductRpcService",
  volumeThreshold: 5,
  rollingCountTimeout: 10000,
  rollingCountBuckets: 10,
};