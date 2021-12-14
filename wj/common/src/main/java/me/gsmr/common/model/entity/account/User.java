package me.gsmr.common.model.entity.account;

import lombok.Data;

@Data
public class User {

    private Long id;

    private String name;

    private String username;

    private String email;

    private String passwordHash;
}
