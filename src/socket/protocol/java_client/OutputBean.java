package  beans;

/**
 * Created by zys on 2017/1/4 0004.
 */
public class Output {
    /**
     * msgType : send SMS
     */

    private MetaBean meta;
    /**
     * meta : {"msgType":"send SMS"}
     * content : hahaha
     */

    private String content;

    public MetaBean getMeta() {
        return meta;
    }

    public void setMeta(MetaBean meta) {
        this.meta = meta;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public static class MetaBean {
        private String msgType;

        public String getMsgType() {
            return msgType;
        }

        public void setMsgType(String msgType) {
            this.msgType = msgType;
        }
    }
}
