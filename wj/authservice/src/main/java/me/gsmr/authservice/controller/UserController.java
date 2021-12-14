package me.gsmr.authservice.controller;

import me.gsmr.common.model.dto.account.UserDto;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/users")
public class UserController {

    @Autowired
    Environment environment;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<UserDto> index() {
        List<UserDto> dtos = new ArrayList<>();
        dtos.add(UserDto.builder().username("name1").username(environment.getProperty("server.port")).build());
        dtos.add(UserDto.builder().username("name2").build());
        return dtos;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto create(@RequestBody UserDto userDto) {
        return null;
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto show(@PathVariable(value = "id") long id) {
        return null;
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto update(@PathVariable long id, @RequestBody UserDto userDto) {
        return null;
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        return null;
    }
}
