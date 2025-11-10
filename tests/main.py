"""Cider API Tests"""
import logging
from client import CiderAPI

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)

logger = logging.getLogger(__name__)


def auth_test():
    api = CiderAPI()

    USERNAME = 'admin'
    PASSWORD = '12345678'


    res = api.register(
        username=USERNAME,
        password=PASSWORD,
    )

    if res.error is not None:
        logger.critical("Got error: %s", res.error)
        return
    
    logger.info("success register: ID - \"%s\"", res.user_id)

    res = api.login(
        username=USERNAME,
        password=PASSWORD,
    )

    if res.error is not None:
        logger.critical("Got error: %s", res.error)
        return
    logger.info("success auth: token - \"%s\"", res.token)

    token = res.token

    res = api.validate(
        token=token,
    )

    if res.error is not None:
        logger.critical("Got error: %s", res.error)
        return
    
    logger.info("success validate: user_id - \"%s\"", res.session.user_id)


def main():
    logger.info("="*70)
    logger.info("Starting Cider API Test")
    logger.info("="*70)
    
    auth_test()
    
    logger.info("="*70)
    logger.info("Test completed")
    logger.info("="*70)


if __name__ == '__main__':
    main()