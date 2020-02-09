class ApplicationController < ActionController::API
  def hello
    greeter_stub = Greeter::Greeter::Stub.new(
      'localhost:5300',
      :this_channel_is_insecure
    )
    req = Greeter::HelloRequest.new(name: 'kohei')
    res = greeter_stub.say_hello(req)
    render json: { message: res.message }
  end
end
