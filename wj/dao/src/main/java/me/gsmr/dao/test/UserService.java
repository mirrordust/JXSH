package me.gsmr.dao.test;




import me.gsmr.common.model.entity.account.User;

import java.util.List;

public interface UserService {
    List<User> findAll();

    User findById(long id);
}
