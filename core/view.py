#!/usr/bin/env python

from .utils import Singleton
from loguru import logger

class CoreView(metaclass=Singleton):
    def __init__(self):
        logger.info("View imported")

