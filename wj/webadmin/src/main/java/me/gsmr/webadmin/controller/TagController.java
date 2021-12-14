package me.gsmr.webadmin.controller;

import me.gsmr.common.model.dto.post.TagDto;
import me.gsmr.webadmin.service.PostServiceClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/tags")
public class TagController {

    @Autowired
    private PostServiceClient client;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<TagDto> index() {
        return client.indexTag();
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public TagDto create(@RequestBody TagDto tagDto) {
        return client.createTag(tagDto);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public TagDto show(@PathVariable long id) {
        return client.showTag(id);
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public TagDto update(@PathVariable long id, @RequestBody TagDto tagDto) {
        return client.updateTag(id, tagDto);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        return client.deleteTag(id);
    }
}
