package me.gsmr.entity.post;

import lombok.Data;

import java.io.Serializable;
import java.sql.Date;

@Data
public class Post implements Serializable {
    private static final long serialVersionUID = 1L;

    private Long id;
    private String title;
    private String body;
//    private String renderedBody;
//    private Boolean published;
//    private Date publishedAt;
//    private String viewName;
//    private Long views;
}
