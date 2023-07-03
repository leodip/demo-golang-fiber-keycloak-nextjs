import Cryptr from "cryptr";

export function encrypt(text) {
  const secretKey = process.env.NEXTAUTH_SECRET;
  const cryptr = new Cryptr(secretKey);

  const encryptedString = cryptr.encrypt(text);
  return encryptedString; 
}

export function decrypt(encryptedString) {
    const secretKey = process.env.NEXTAUTH_SECRET;
    const cryptr = new Cryptr(secretKey);
  
    const text = cryptr.decrypt(encryptedString);
    return text;
  }