# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: isuxportal/resources/staff.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("isuxportal/resources/staff.proto", :syntax => :proto3) do
    add_message "isuxportal.proto.resources.Staff" do
      optional :id, :int64, 1
      optional :github_login, :string, 2
    end
  end
end

module Isuxportal
  module Proto
    module Resources
      Staff = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.resources.Staff").msgclass
    end
  end
end
