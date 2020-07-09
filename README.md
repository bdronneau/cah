# CAH - Card Against Humanity

[![Build Status](https://travis-ci.com/bdronneau/cah.svg?branch=master)](https://travis-ci.com/bdronneau/cah) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_cah&metric=alert_status)](https://sonarcloud.io/dashboard?id=bdronneau_cah) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_cah&metric=code_smells)](https://sonarcloud.io/dashboard?id=bdronneau_cah) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_cah&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=bdronneau_cah) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_cah&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=bdronneau_cah) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_cah&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=bdronneau_cah)

I develop this open source game based on [Cards against Humanity](https://cardsagainsthumanity.com/) during COVID in order to try golang and vuejs together !

## Status
- Game working
- Code does not have unit tests
- I'm using `setInterval` for refresh views, `websocket` will me more efficient
- Tables are in schema `public`
- vuex is not very well implemented

Despite these behaviours, I'm pretty happy because it's working and I can play with my friends.
Moreover, I learn a lot on vuejs development and golang development ...

If I give more love on this project ?
- Use a better structure for Golang API
- Add context logging
- Add prometheus metrics
- Limit the SPA scope to game part and make golang templates for lobby, home, room creation ...
- Include unit tests and cypress

## Usage
- [front](./front/README.md)
- [api](./api/README.md)

## Hosting
- [scripts/cah.yml](./scripts/cah.yml) contain a example with dummy password for deployment compatible with [deploy](https://github.com/ViBiOh/deploy) and [traefik](https://containo.us/traefik/)

## Credits

Sources for cards
- https://github.com/crhallberg/json-against-humanity
- https://github.com/lhoang/Cards-Against-Humanity---Generator


## License

This project is licensed under the MIT license (see LICENSE file).
