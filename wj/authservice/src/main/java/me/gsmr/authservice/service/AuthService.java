package me.gsmr.authservice.service;

import me.gsmr.common.model.dto.account.CredentialDto;
import me.gsmr.common.model.dto.account.UserDto;

public interface AuthService {

    CredentialDto authenticateByEmailPassword(UserDto userDto);
}
