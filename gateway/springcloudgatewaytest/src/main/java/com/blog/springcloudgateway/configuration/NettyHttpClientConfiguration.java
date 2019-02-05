package com.blog.springcloudgateway.configuration;

import org.springframework.context.annotation.Configuration;

// import org.springframework.beans.factory.annotation.Qualifier;
// import org.springframework.context.annotation.Bean;
// import org.springframework.context.annotation.Configuration;

// import io.netty.util.internal.shaded.org.jctools.queues.MessagePassingQueue.Consumer;
// import reactor.netty.http.client.HttpClient;

@Configuration
class NettyHttpClientConfiguration {
    // @Bean
    // HttpClient httpClient(@Qualifier("nettyClientOptions") Consumer<? super HttpClientOptions.Builder> options) {
    //     return HttpClient.create(options.andThen(this::enableCompressionSupport));
    // }

    // private void enableCompressionSupport(Object builder) {
    //     ((HttpClientOptions.Builder) builder).compression(true);
    // }
}