package me.gsmr.dao;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@MapperScan("me.gsmr.dao.mapper")
@SpringBootApplication
public class DaoApplication {

    public static void main(String[] args) {
        SpringApplication.run(DaoApplication.class, args);
    }

}
