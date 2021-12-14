package me.gsmr.common.model.dto.post;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class PostDto implements Serializable {
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
