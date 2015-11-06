#!/usr/bin/env python

import sys
import time

from faker import Factory


fake = Factory.create()

while True:
    try:
        sys.stdout.write(fake.sentence())
        sys.stdout.write('\n')
        sys.stdout.flush()
        time.sleep(0.1)
    except KeyboardInterrupt:
        sys.exit(0)
