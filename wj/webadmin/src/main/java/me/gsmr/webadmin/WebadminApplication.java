package me.gsmr.webadmin;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@EnableFeignClients
@SpringBootApplication
public class WebadminApplication {

    public static void main(String[] args) {
        SpringApplication.run(WebadminApplication.class, args);
    }

}
