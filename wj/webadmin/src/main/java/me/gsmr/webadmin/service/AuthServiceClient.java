package me.gsmr.webadmin.service;

import me.gsmr.common.model.dto.account.UserDto;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@FeignClient("w-authservice")
public interface AuthServiceClient {

    /* User */
    @GetMapping("/users")
    List<UserDto> indexUser();

    @PostMapping("/users")
    UserDto createUser(@RequestBody UserDto userDto);

    @GetMapping("/users/{id}")
    UserDto showUser(@PathVariable long id);

    @PutMapping("/users/{id}")
    UserDto updateUser(@PathVariable long id, @RequestBody UserDto userDto);

    @DeleteMapping("/users/{id}")
    Object deleteUser(@PathVariable long id);
}
