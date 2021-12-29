package me.gsmr.authservice.service.impl;

import me.gsmr.authservice.exception.DbException;
import me.gsmr.authservice.mapper.UserMapper;
import me.gsmr.authservice.service.AuthService;
import me.gsmr.common.model.dto.account.CredentialDto;
import me.gsmr.common.model.dto.account.UserDto;
import me.gsmr.common.model.entity.account.User;
import me.gsmr.common.utils.Enc;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service
public class AuthServiceImpl implements AuthService {

    @Resource
    private UserMapper userMapper;

    @Override
    public CredentialDto authenticateByEmailPassword(UserDto userDto) {
        String decrypted = Enc.decrypt(userDto.getPassword());
        userDto.setPasswordHash(Enc.hash(decrypted));
        User user;
        try {
            user = userMapper.findUserByEmailPasswordHash(userDto.getEmail(), userDto.getPasswordHash());
        } catch (Exception e) {
            throw new DbException("db select error", e);
        }
        if (user == null) {
            return CredentialDto.builder().authenticated(false).build();
        }
        return CredentialDto.builder().authenticated(true).accessToken("token_dev").build();
    }
}
