import unittest
from __init__ import get_ip

class TestGetIP(unittest.TestCase):
    def test_get_result(self):
        ip = get_ip()
        self.assertNotEqual(len(ip), 0, "got something")

if __name__ == '__main__':
    unittest.main()
