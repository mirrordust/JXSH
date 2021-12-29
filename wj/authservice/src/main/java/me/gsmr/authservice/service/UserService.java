package me.gsmr.authservice.service;

import me.gsmr.common.model.dto.account.UserDto;

import java.util.List;

public interface UserService {

    List<UserDto> findAllUsers();

    UserDto findUserById(long id);

    UserDto createUser(UserDto userDto);

    UserDto updateUser(long id, UserDto userDto);

    int deleteUser(long id);
}
