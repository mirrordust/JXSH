package me.gsmr.common.model.dto.account;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class CredentialDto implements Serializable {
    private static final long serialVersionUID = 1L;

    private String accessToken;

    private Boolean authenticated;
}
