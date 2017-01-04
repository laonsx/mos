//java对应的封包过程

    public static byte[] packet(String message) {
        byte[] re;
    	try {
    		byte[] s=message.getBytes("utf-8");
    		System.out.println(message.length());
    		byte[] a=intToByteArray1(s.length);
    		byte[] b="Headers".getBytes("utf-8");
    		re=new byte[s.length+a.length+b.length];
    		System.arraycopy(b, 0, re, 0, b.length);
    		System.arraycopy(a, 0, re, b.length, a.length);
    		System.arraycopy(s, 0, re, a.length+b.length, s.length);
    		return re;
    	} catch (UnsupportedEncodingException e) {
    		e.printStackTrace();
    		return null;
    	}
    }

    public static byte[] intToByteArray1(int i) {
        byte[] result = new byte[4];
        result[0] = (byte)((i >> 24) & 0xFF);
        result[1] = (byte)((i >> 16) & 0xFF);
        result[2] = (byte)((i >> 8) & 0xFF);
        result[3] = (byte)(i & 0xFF);
        return result;
    }
