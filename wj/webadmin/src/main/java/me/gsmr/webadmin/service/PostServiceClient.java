package me.gsmr.webadmin.service;

import me.gsmr.common.model.dto.post.CategoryDto;
import me.gsmr.common.model.dto.post.PostDto;
import me.gsmr.common.model.dto.post.TagDto;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@FeignClient("w-postservice")
public interface PostServiceClient {

    /* Post */
    @GetMapping("/posts")
    List<PostDto> indexPost();

    @PostMapping("/posts")
    PostDto createPost(@RequestBody PostDto postDto);

    @GetMapping("/posts/{id}")
    PostDto showPost(@PathVariable long id);

    @PutMapping("/posts/{id}")
    PostDto updatePost(@PathVariable long id, @RequestBody PostDto postDto);

    @DeleteMapping("/posts/{id}")
    Object deletePost(@PathVariable long id);

    /* Tag */
    @GetMapping("/tags")
    List<TagDto> indexTag();

    @PostMapping("/tags")
    TagDto createTag(@RequestBody TagDto tagDto);

    @GetMapping("/tags/{id}")
    TagDto showTag(@PathVariable long id);

    @PutMapping("/tags/{id}")
    TagDto updateTag(@PathVariable long id, @RequestBody TagDto tagDto);

    @DeleteMapping("/tags/{id}")
    Object deleteTag(@PathVariable long id);

    /* Category */
    @GetMapping("/categories")
    List<CategoryDto> indexCategory();

    @PostMapping("/categories")
    CategoryDto createCategory(@RequestBody CategoryDto categoryDto);

    @GetMapping("/categories/{id}")
    CategoryDto showCategory(@PathVariable long id);

    @PutMapping("/categories/{id}")
    CategoryDto updateCategory(@PathVariable long id, @RequestBody CategoryDto categoryDto);

    @DeleteMapping("/categories/{id}")
    Object deleteCategory(@PathVariable long id);
}
