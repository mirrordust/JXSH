package me.gsmr.authservice.mapper;

import me.gsmr.common.model.entity.account.User;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;

@Mapper
public interface UserMapper {

    List<User> findAllUsers();

    User findUserById(long id);

    int createUser(User user);

    int updateUser(User user);

    int deleteUser(long id);

    User findUserByEmailPasswordHash(@Param("email") String email, @Param("passwordHash") String passwordHash);
}
