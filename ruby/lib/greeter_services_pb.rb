# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: greeter.proto for package 'greeter'

require 'grpc'
require 'greeter_pb'

module Greeter
  module Greeter
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'greeter.Greeter'

      rpc :SayHello, HelloRequest, HelloReply
    end

    Stub = Service.rpc_stub_class
  end
end
