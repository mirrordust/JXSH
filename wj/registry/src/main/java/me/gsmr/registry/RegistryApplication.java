package me.gsmr.registry;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.server.EnableEurekaServer;

@EnableEurekaServer
@SpringBootApplication
public class RegistryApplication {

    public static void main(String[] args) {
        SpringApplication.run(RegistryApplication.class, args);
    }

    /* References:
    1. https://medium.com/swlh/spring-cloud-high-availability-for-eureka-b5b7abcefb32
        need to modify hosts with:
        127.0.0.1 peer-1-server.com
        127.0.0.1 peer-2-server.com
        127.0.0.1 peer-3-server.com
    */

}
