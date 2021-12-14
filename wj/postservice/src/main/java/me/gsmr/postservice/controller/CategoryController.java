package me.gsmr.postservice.controller;

import me.gsmr.common.model.dto.post.CategoryDto;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/categories")
public class CategoryController {

    @Autowired
    Environment environment;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<CategoryDto> index() {
        List<CategoryDto> dtos = new ArrayList<>();
        dtos.add(new CategoryDto(1L, environment.getProperty("server.port")));
        dtos.add(new CategoryDto(2L, "n2"));
        return dtos;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CategoryDto create(@RequestBody CategoryDto categoryDto) {
        return null;
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public CategoryDto show(@PathVariable long id) {
        return null;
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public CategoryDto update(@PathVariable long id, @RequestBody CategoryDto categoryDto) {
        return null;
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        return null;
    }
}
