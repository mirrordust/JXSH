package me.gsmr.entity.account;

import lombok.Data;

import java.io.Serializable;

@Data
public class User implements Serializable {
    private static final long serialVersionUID = 1L;

    private Long id;
    private String name;
    private String username;
    private String email;
    private String password;
}
