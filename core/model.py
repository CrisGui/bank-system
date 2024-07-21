#!/usr/bin/env python

from .utils import Singleton

from loguru import logger

import os
# import mysql.connector

class CoreModel(metaclass=Singleton):
    def __init__(self):
        logger.info("Model imported")

        # try:
        #     connection = mysql.connector.connect(user=os.getenv("MYSQL_USER"), password=os.getenv("MYSQL_ROOT_PASSWORD"),host="localhost",databse="bank")
        #     if connection.is_connected():
        #         logger.debug("Database connected")
        # except mysql.connector.Error as e:
        #     logger.critical(f"Error: {e}")
        # finally:
        #     if "connection" in locals():
        #         connection.close()
        #         logger.debug("Database disconnected")
