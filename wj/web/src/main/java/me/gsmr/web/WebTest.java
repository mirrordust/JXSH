package me.gsmr.web;

import me.gsmr.dao.DaoTest;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("test")
public class WebTest {
    @RequestMapping("showAll")
    public String showAll() {
        DaoTest daoTest = new DaoTest();
        return daoTest.showDao() + "  Web! ";
    }
}
