package me.gsmr.webadmin.controller;

import me.gsmr.common.model.dto.post.PostDto;
import me.gsmr.webadmin.service.PostServiceClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/posts")
public class PostController {

    @Autowired
    private PostServiceClient client;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<PostDto> index() {
        return client.indexPost();
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public PostDto create(@RequestBody PostDto postDto) {
        return client.createPost(postDto);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public PostDto show(@PathVariable long id) {
        return client.showPost(id);
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public PostDto update(@PathVariable long id, @RequestBody PostDto postDto) {
        return client.updatePost(id, postDto);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        return client.deletePost(id);
    }
}
