package me.gsmr.webadmin.controller;

import me.gsmr.common.model.dto.account.UserDto;
import me.gsmr.webadmin.service.AuthServiceClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/users")
public class UserController {

    private final Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    private AuthServiceClient client;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<UserDto> index(@RequestHeader HttpHeaders headers) {
        logger.info(headers.toString());
        return client.indexUser();
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto create(@RequestBody UserDto userDto) {
        return client.createUser(userDto);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto show(@PathVariable(value = "id") long id) {
        return client.showUser(id);
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto update(@PathVariable long id, @RequestBody UserDto userDto) {
        return client.updateUser(id, userDto);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object delete(@PathVariable long id) {
        return client.deleteUser(id);
    }
}
