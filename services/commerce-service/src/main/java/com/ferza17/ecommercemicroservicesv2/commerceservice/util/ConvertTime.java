package com.ferza17.ecommercemicroservicesv2.commerceservice.util;

import com.google.protobuf.Timestamp;

import java.time.Instant;

public class ConvertTime {
    public static Instant toInstant(Timestamp ts) {
        return Instant.ofEpochSecond(ts.getSeconds(), ts.getNanos());
    }

    public static Timestamp toTimestamp(Instant instant) {
        if (instant == null) {
            return Timestamp.getDefaultInstance();
        }
        return Timestamp.newBuilder()
                .setSeconds(instant.getEpochSecond())
                .setNanos(instant.getNano())
                .build();
    }
}
