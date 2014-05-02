'use strict';

describe('Service: ServicesJs', function () {

  // load the service's module
  beforeEach(module('miranalApp'));

  // instantiate service
  var ServicesJs;
  beforeEach(inject(function (_ServicesJs_) {
    ServicesJs = _ServicesJs_;
  }));

  it('should do something', function () {
    expect(!!ServicesJs).toBe(true);
  });

});
