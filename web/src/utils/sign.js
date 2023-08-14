import {secretKey} from '@/config/index.js'
import md5 from 'crypto-js/md5';
import Base64 from 'crypto-js/enc-base64';
import hmacSHA256 from 'crypto-js/hmac-sha256';

export function paramsSignMd5(params, now, crypto="md5") {
    const sortData = { } // 排序后的对象
    Object.keys(params).sort().map(key => {
        sortData[key] = params[key]
    })
    var needSignatureStr = now + "|";
    for (var key in sortData) {
        var value = sortData[key];
        var type = Object.prototype.toString.call(value)
        if (type === '[object Object]' || type === '[object Array]') {
            value = JSON.stringify(value)
        }
        needSignatureStr = needSignatureStr + key + '=' + value + '&';
    }
    
    if(crypto == "sha256") {
        return Base64.stringify(hmacSHA256(needSignatureStr.substring(0, needSignatureStr.length - 1), secretKey));
    }
    needSignatureStr += 'key=' + secretKey;
    console.log('sign:', needSignatureStr);
    return Base64.stringify(md5(needSignatureStr));
}