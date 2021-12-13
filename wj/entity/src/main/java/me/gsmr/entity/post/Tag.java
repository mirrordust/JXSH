package me.gsmr.entity.post;

import lombok.Data;

import java.io.Serializable;

@Data
public class Tag implements Serializable {
    private static final long serialVersionUID = 1L;

    private Long id;
    private String name;
}
