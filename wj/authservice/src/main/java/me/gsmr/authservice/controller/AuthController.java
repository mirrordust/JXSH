package me.gsmr.authservice.controller;

import me.gsmr.authservice.service.AuthService;
import me.gsmr.common.model.dto.account.CredentialDto;
import me.gsmr.common.model.dto.account.UserDto;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

@RestController
@RequestMapping("/auth")
public class AuthController {

    @Resource
    private AuthService authService;

    @PostMapping("/sessions")
    @ResponseStatus(HttpStatus.CREATED)
    public CredentialDto login(@RequestBody UserDto userDto) {
        return authService.authenticateByEmailPassword(userDto);
    }

    @DeleteMapping("/sessions")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public Object logout(@RequestBody CredentialDto credentialDto) {
        return null;
    }
}
