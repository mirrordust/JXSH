package me.gsmr.dao.test;

import me.gsmr.dao.mapper.UserMapper;
import me.gsmr.entity.account.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private UserMapper userMapper;

    @Override
    public List<User> findAll() {
        return userMapper.findAll();
    }

    @Override
    public User findById(long id) {
        return userMapper.findById(id);
    }
}
