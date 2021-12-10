package me.gsmr.dao;

import me.gsmr.entity.EntityTest;

public class DaoTest {

    public String showDao() {
        EntityTest entityTest = new EntityTest();
        return entityTest.showEntity() + " Dao! ";
    }
}
