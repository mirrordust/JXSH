package me.gsmr.dao.mapper;

import me.gsmr.entity.account.User;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;

@Mapper
public interface UserMapper {
    List<User> findAll();

    User findById(long id);
}
