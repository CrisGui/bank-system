#!/usr/bin/env python

from .utils import Singleton
from loguru import logger

class Model(metaclass=Singleton):
    def __init__(self):
        logger.info("Model imported")

