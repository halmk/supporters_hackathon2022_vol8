FROM ubuntu:20.04

# Install Node.js, Yarn and required dependencies
RUN apt-get update \
  && apt-get install -y curl gnupg build-essential \
  && curl -fsSL https://deb.nodesource.com/setup_16.x | bash - \
  && curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
  && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list \
  && apt-get remove -y --purge cmdtest \
  && apt-get update \
  && apt-get install -y locales \
  && locale-gen ja_JP.UTF-8 \
  && echo "export LANG=ja_JP.UTF-8" >> ~/.bashrc \
  && apt-get install -y nodejs yarn g++ make tmux vim git \
  # remove useless files from the current layer
  && rm -rf /var/lib/apt/lists/* \
  && rm -rf /var/lib/apt/lists.d/* \
  && apt-get autoremove \
  && apt-get clean \
  && apt-get autoclean

