package me.gsmr.authservice.controller;

import me.gsmr.authservice.service.UserService;
import me.gsmr.common.model.dto.account.UserDto;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.server.ResponseStatusException;

import javax.annotation.Resource;
import java.util.List;

@RestController
@RequestMapping("/users")
public class UserController {

    private final Logger logger = LoggerFactory.getLogger(getClass());

    @Resource
    private UserService userService;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<UserDto> index() {
        return userService.findAllUsers();
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto show(@PathVariable(value = "id") long id) {
        UserDto userDto = userService.findUserById(id);
        if (userDto == null) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "no such user");
        }
        return userDto;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto create(@RequestBody UserDto userDto) {
        UserDto userDto1 = userService.createUser(userDto);
        if (userDto == null) {
            throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "create fail");
        }
        return userDto1;
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto update(@PathVariable long id, @RequestBody UserDto userDto) {
        UserDto userDto1 = userService.updateUser(id, userDto);
        if (userDto1 == null) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "no such user");
        }
        return userDto1;
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        int rows = userService.deleteUser(id);
        if (rows == 0) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "no such user");
        }
        return null;
    }
}
