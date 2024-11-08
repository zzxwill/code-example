package example;

import com.volcengine.vefaas.server.Server;

public class Main {
    public static void main(String[] args) throws Exception {
        new Server(StreamHandler.class.newInstance());
    }
}