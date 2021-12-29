package me.gsmr.authservice.service.impl;

import me.gsmr.authservice.exception.DbException;
import me.gsmr.authservice.mapper.UserMapper;
import me.gsmr.authservice.service.UserService;
import me.gsmr.common.model.dto.account.UserDto;
import me.gsmr.common.model.entity.account.User;
import me.gsmr.common.utils.Enc;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class UserServiceImpl implements UserService {

    @Resource
    private UserMapper userMapper;

    @Override
    public List<UserDto> findAllUsers() {
        List<User> users;
        try {
            users = userMapper.findAllUsers();
        } catch (Exception e) {
            throw new DbException("db select error", e);
        }

        if (users == null) {
            return null;
        }
        return users.stream().map(this::transform0).collect(Collectors.toList());
    }

    @Override
    public UserDto findUserById(long id) {
        User user;
        try {
            user = userMapper.findUserById(id);
        } catch (Exception e) {
            throw new DbException("db select error", e);
        }

        if (user == null) {
            return null;
        }
        return transform0(user);
    }

    @Override
    public UserDto createUser(UserDto userDto) {
        User user = transform1(userDto);
        int rows;
        try {
            rows = userMapper.createUser(user);
        } catch (Exception e) {
            throw new DbException("db create error", e);
        }

        if (rows == 1) {
            return transform0(user);
        } else {
            return null;
        }
    }

    @Override
    public UserDto updateUser(long id, UserDto userDto) {
        userDto.setId(id);
        User user = transform1(userDto);
        int rows;
        try {
            rows = userMapper.updateUser(user);
        } catch (Exception e) {
            throw new DbException("db update error", e);
        }

        if (rows == 1) {
            return transform0(user);
        } else {
            return null;
        }
    }

    @Override
    public int deleteUser(long id) {
        int rows;
        try {
            rows = userMapper.deleteUser(id);
        } catch (Exception e) {
            throw new DbException("db delete error", e);
        }

        return rows;
    }

    /**
     * User -> UserDto
     */
    private UserDto transform0(User user) {
        if (user == null) {
            return null;
        }
        return UserDto.builder().id(user.getId()).name(user.getName()).username(user.getUsername()).email(user.getEmail()).passwordHash(user.getPasswordHash()).build();
    }

    /**
     * UserDto -> User
     */
    private User transform1(UserDto userDto) {
        if (userDto == null) {
            return null;
        }
        User user = new User();
        user.setId(userDto.getId());
        user.setName(userDto.getName());
        user.setUsername(userDto.getUsername());
        user.setEmail(userDto.getEmail());
        user.setPasswordHash(Enc.hash(userDto.getPassword()));
        return user;
    }

}
