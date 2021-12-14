package me.gsmr.webadmin.controller;

import me.gsmr.common.model.dto.post.CategoryDto;
import me.gsmr.webadmin.service.PostServiceClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/categories")
public class CategoryController {

    @Autowired
    private PostServiceClient client;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<CategoryDto> index() {
        return client.indexCategory();
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CategoryDto create(@RequestBody CategoryDto categoryDto) {
        return client.createCategory(categoryDto);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public CategoryDto show(@PathVariable long id) {
        return client.showCategory(id);
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public CategoryDto update(@PathVariable long id, @RequestBody CategoryDto categoryDto) {
        return client.updateCategory(id, categoryDto);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        return client.deleteCategory(id);
    }
}
