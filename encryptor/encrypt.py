import logging
import os
import sys
import binascii
from Crypto.Cipher import AES


MODE = AES.MODE_CFB
BLOCK_SIZE = 16
SEGMENT_SIZE = 128

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG, format='%(message)s')


class Encryptor:
    @classmethod
    def _pad_string(cls, value):
        length = len(value)
        pad_size = BLOCK_SIZE - (length % BLOCK_SIZE)
        return (value.ljust(length + pad_size, '\x00')).encode('utf-8')

    @classmethod
    def encrypt(cls, key, iv, plaintext):
        key_bytes = bytes(key.encode('utf-8'))
        iv_bytes = bytes(iv.encode('utf-8'))
        aes = AES.new(key_bytes, MODE, iv_bytes, segment_size=SEGMENT_SIZE)
        plaintext = cls._pad_string(plaintext)
        encrypted_text = aes.encrypt(plaintext)
        return binascii.hexlify(encrypted_text)


def main():
    if len(sys.argv) != 2:
        raise ValueError("No plain text provided")

    plain_text = sys.argv[1]
    token = Encryptor.encrypt(
        os.getenv('PROXY_AES_KEY'),
        os.getenv('PROXY_AES_IV'),
        plain_text)
    encrypted = token.decode('utf-8')
    return encrypted


if __name__ == '__main__':
    logger.info(main())
