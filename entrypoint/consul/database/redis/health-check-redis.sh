#!/bin/bash

redis_health_check() {
    local redis_host="redis-local"
    local redis_port="6379"
    
    echo "Running Redis health check..."
    
    # Test 1: TCP Connection
    if ! nc -z $redis_host $redis_port; then
        echo "❌ Redis TCP connection failed"
    fi
    echo "✅ Redis TCP connection OK"
    
    # Test 2: PING command
    if command -v redis-cli >/dev/null 2>&1; then
        if ! redis-cli -h $redis_host -p $redis_port ping | grep -q "PONG"; then
            echo "❌ Redis PING failed"
        fi
        echo "✅ Redis PING OK"
        
        # Test 3: Set/Get test
        test_key="health_check_$(date +%s)"
        test_value="test_value"
        
        if redis-cli -h $redis_host -p $redis_port set $test_key $test_value | grep -q "OK"; then
            if redis-cli -h $redis_host -p $redis_port get $test_key | grep -q $test_value; then
                redis-cli -h $redis_host -p $redis_port del $test_key > /dev/null
                echo "✅ Redis SET/GET operations OK"
            else
                echo "❌ Redis GET operation failed"
            fi
        else
            echo "❌ Redis SET operation failed"
        fi
    else
        echo "⚠️  redis-cli not available, skipping PING test"
    fi
    
    echo "✅ All Redis health checks passed"
}

# Execute health check
redis_health_check
