package me.gsmr.common.utils;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public class Enc {

    public static String encrypt(String plaintext) {
        return plaintext;
    }

    public static String decrypt(String ciphertext) {
        return ciphertext;
    }

    public static String hash(String plaintext) {
        return plaintext;
//        byte[] bytes = plaintext.getBytes(StandardCharsets.UTF_8);
//
//        MessageDigest md;
//        try {
//            md = MessageDigest.getInstance("MD5");
//        } catch (NoSuchAlgorithmException e) {
//            e.printStackTrace();
//            return "";
//        }
//
//        byte[] md5Digest = md.digest(bytes);
//
//        return new String(md5Digest, StandardCharsets.UTF_8);
    }

}
