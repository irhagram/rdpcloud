<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/sysinfo/time.proto

namespace GPBMetadata\Services\Sysinfo;

class Time
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
services/sysinfo/time.protoservices.sysinfo"
GetUptimeRequest"#
GetUptimeResponse
uptime (B5Z3github.com/s77rt/rdpcloud/proto/go/services/sysinfobproto3'
        , true);

        static::$is_initialized = true;
    }
}

