package me.gsmr.web;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.discovery.DiscoveryClient;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RefreshScope
@RestController
@RequestMapping("test")
public class WebTest {

    private String name;

    @Autowired
    private DiscoveryClient discoveryClient;

    @Value("${message:Hello default}")
    private String message;

    @RequestMapping("showAll")
    public String showAll() {
        return " From Web! " + message;
    }

    @RequestMapping("/service-instance/{applicationName}")
    public List<ServiceInstance> serviceInstancesByApplicationName(@PathVariable String applicationName) {
        return this.discoveryClient.getInstances(applicationName);
    }

}
