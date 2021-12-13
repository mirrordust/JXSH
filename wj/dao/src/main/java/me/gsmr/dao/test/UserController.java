package me.gsmr.dao.test;

import me.gsmr.entity.account.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/users")
public class UserController {

    @Autowired
    private UserService userService;

    @RequestMapping("/findAll")
    public List<User> findAll() {
        return userService.findAll();
    }

    @RequestMapping("/findById")
    public User findById(@RequestParam("id") long id) {
        return userService.findById(id);
    }
}
